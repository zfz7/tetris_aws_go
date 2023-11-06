import com.github.gradle.node.NodeExtension
import com.github.gradle.node.yarn.task.YarnTask

plugins {
    id("com.github.node-gradle.node")
}

configure<NodeExtension> {
    version.set(rootProject.extra["nodeVersion"] as String)
    yarnVersion.set(rootProject.extra["yarnVersion"] as String)
    download.set(true)
}

val install = tasks.register<YarnTask>("install") {
    dependsOn("copy_ts_client_from_model")
    inputs.file(file("$projectDir/yarn.lock"))
    inputs.file(file("$projectDir/package.json"))
    outputs.dir(file("$projectDir/node_modules"))
    args.set(listOf("install"))
}

tasks.register<Copy>("copy_ts_client_from_model") {
    dependsOn(":model:build")
    from(file("${project(":model").projectDir}/build/smithyprojections/model/source/typescript-codegen"))
    into(file("$projectDir"))
}

tasks.register<YarnTask>("test") {
    environment.set(mapOf("CI" to "true"))
    dependsOn(install)
    args.set(listOf("test"))
}

tasks.register<YarnTask>("build") {
    dependsOn(install)
    mustRunAfter("test")
    inputs.dir(file("$projectDir"))
    outputs.dir(file("$projectDir/build"))
    args.set(listOf("build"))
}

task<Delete>("clean") {
    delete(fileTree("$projectDir") {
        exclude("yarn.lock", "build.gradle.kts", "settings.gradle.kts", "/node_modules/**")
    })
}