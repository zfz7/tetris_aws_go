import { Construct } from 'constructs';
import { Duration, Stack, StackProps } from 'aws-cdk-lib';
import { Bucket, BucketAccessControl } from 'aws-cdk-lib/aws-s3';
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment';
import * as path from 'path';
import { AccessLevel, Distribution, PriceClass, ViewerProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { ARecord, IPublicHostedZone, RecordTarget } from 'aws-cdk-lib/aws-route53';
import { DnsValidatedCertificate } from 'aws-cdk-lib/aws-certificatemanager';
import { CloudFrontTarget } from 'aws-cdk-lib/aws-route53-targets';
import { S3BucketOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';
import { PROJECT } from '../config';

export interface WebsiteStackProps extends StackProps {
  hostedZone: IPublicHostedZone;
  domainName: string;
}

export class WebsiteStack extends Stack {
  constructor(scope: Construct, id: string, props: WebsiteStackProps) {
    super(scope, id, props);

    const websiteBuck = new Bucket(this, `${PROJECT}-Website-Bucket`, {
      accessControl: BucketAccessControl.PRIVATE,
    });

    //Should use DnsValidatedCertificate per:
    //https://docs.aws.amazon.com/cdk/api/v1/docs/aws-certificatemanager-readme.html#cross-region-certificates
    const certificate = new DnsValidatedCertificate(this, `${PROJECT}-Website-Certificate`, {
      hostedZone: props.hostedZone,
      domainName: props.domainName,
      region: 'us-east-1',
    });

    const distribution = new Distribution(this, `${PROJECT}-Distribution`, {
      defaultRootObject: 'index.html',
      comment: props.domainName,
      defaultBehavior: {
        origin: S3BucketOrigin.withOriginAccessControl(websiteBuck, {
          originAccessLevels: [AccessLevel.READ, AccessLevel.LIST],
        }),
        viewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
      },
      domainNames: [props.domainName],
      certificate: certificate,
      priceClass: PriceClass.PRICE_CLASS_100, // USA, Canada, Europe, & Israel
      errorResponses: [
        //SPA applications handle routing on the client side
        {
          httpStatus: 404, // Catch missing routes and serve application
          responseHttpStatus: 200,
          responsePagePath: '/index.html', // Serve index.html for missing routes then
          ttl: Duration.seconds(0), // Cache 404 responses for 0 seconds
        },
      ],
    });

    new BucketDeployment(this, `${PROJECT}-Bucket-Deployment`, {
      destinationBucket: websiteBuck,
      distribution: distribution,
      distributionPaths: ['/*'], // Invalidate existing distribution
      sources: [Source.asset(path.resolve('../frontend/', 'build'))],
    });

    new ARecord(this, `${PROJECT}-Website-ARecord`, {
      zone: props.hostedZone,
      recordName: props.domainName,
      target: RecordTarget.fromAlias(new CloudFrontTarget(distribution)),
    });
  }
}
