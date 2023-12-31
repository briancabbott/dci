# DO NOT EDIT THIS FILE - See: https://eclipse.dev/jetty/documentation/

[description]
Enables Annotation scanning for deployed web applications.

[environment]
ee9

[depend]
ee9-plus

[lib]
lib/jetty-ee9-annotations-${jetty.version}.jar
lib/ee9-annotations/*.jar

[jpms]
add-modules:org.objectweb.asm

