extra["smithyTypeScriptVersion"] = "0.19.0"
extra["smithyVersion"] = "1.42.0"
plugins {
    application
    id("software.amazon.smithy.gradle.smithy-jar").version("0.9.0")
}

repositories {
    mavenLocal()
    mavenCentral()
}

dependencies {
    implementation("software.amazon.smithy.typescript:smithy-typescript-codegen:${rootProject.extra["smithyTypeScriptVersion"]}")
}

smithy {
    smithyBuildConfigs = files("smithy-build.json")
    dependencies {
        implementation("software.amazon.smithy.typescript:smithy-aws-typescript-codegen:${rootProject.extra["smithyTypeScriptVersion"]}")
        implementation("software.amazon.smithy:smithy-openapi:${rootProject.extra["smithyVersion"]}")
        implementation("software.amazon.smithy:smithy-model:${rootProject.extra["smithyVersion"]}")
        implementation("software.amazon.smithy:smithy-aws-traits:${rootProject.extra["smithyVersion"]}")
        implementation("software.amazon.smithy:smithy-aws-apigateway-openapi:${rootProject.extra["smithyVersion"]}")
        implementation("software.amazon.smithy:smithy-cli:${rootProject.extra["smithyVersion"]}")
    }
}
