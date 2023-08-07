#!/bin/bash
set -ex

PROPERTIES_THERE=$(find /vol1/aax2cfg -iname "*.properties")

if [[ "$EXTERNAL_JMS" != "unset" ]]; then
  echo "subsystem.config.jms.namingInitialContextFactory=org.apache.activemq.jndi.ActiveMQInitialContextFactory" > /vol1/aax2cfg/subsystem-interface.properties
  echo "subsystem.config.jms.jndiConnectionFactory=ConnectionFactory" >> /vol1/aax2cfg/subsystem-interface.properties
  echo "subsystem.config.jms.enableAuthentication=false" >> /vol1/aax2cfg/subsystem-interface.properties
  echo "subsystem.config.jms.providerUrl=$EXTERNAL_JMS" >> /vol1/aax2cfg/subsystem-interface.properties
  echo "subsystem.config.localjmsqueuenames={Tremeal: 'dynamicQueues/queue/cs_tremeal'}" >> /vol1/aax2cfg/subsystem-interface.properties
fi

if [[ -z "$AAX_DB_SSL_MODE" || "$AAX_DB_SSL_MODE" == "disable" ]]; then
  sed -i~ "s,###DB_URL###,$AAX_DB_URL," /vol1/aax2cfg/dbconn.properties
else
  SSL_DB_URL="$AAX_DB_URL?sslmode=$AAX_DB_SSL_MODE\&sslcert=$AAX_DB_SSL_CLIENT_CERT_LOCATION\&sslkey=$AAX_DB_SSL_CLIENT_KEY_LOCATION\&sslrootcert=$AAX_DB_SSL_SERVER_CA_LOCATION"
  sed -i~ "s,###DB_URL###,$SSL_DB_URL," /vol1/aax2cfg/dbconn.properties
fi
sed -i~ "s,###DB_USER###,$AAX_DB_USER," /vol1/aax2cfg/dbconn.properties
sed -i~ "s,###DB_PASS###,$AAX_DB_PASS," /vol1/aax2cfg/dbconn.properties

if [[ ! -z "$PROPERTIES_THERE" ]]; then
  for F in /vol1/aax2cfg/*.properties; do
    cp $F $F~
    awk '{ 
        while (match($0, /###[^#]+###/)) { 
            search = substr($0, RSTART + 3, RLENGTH - 6)
            $0 = substr($0, 1, RSTART - 1)   \
                ENVIRON[search]             \
                substr($0, RSTART + RLENGTH) 
        } 
        print 
    }' $F~ > $F
  done
fi

if [[ "$ELASTIC_APM_ENABLED" == "true" ]]; then
	ELASTIC_APM_OPTS="-javaagent:/vol1/elastic-apm-agent.jar -Delastic.apm.service_name=$ELASTIC_APM_SERVICE_NAME -Delastic.apm.application_packages=at.compax -Delastic.apm.server_url=$ELASTIC_APM_URL -Delastic.apm.secret_token=$ELASTIC_APM_SECRET_TOKEN -Delastic.apm.verify_server_cert=$ELASTIC_APM_VERIFY_SERVER_CERT -Delastic.apm.environment=$ELASTIC_APM_ENVIRONMENT -Delastic.apm.transaction_sample_rate=0.1 -Delastic.apm.capture_headers=false"
else
	ELASTIC_APM_OPTS=""
fi

JMX_OPTS="-Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.port=1098 -Dcom.sun.management.jmxremote.rmi.port=1098 -Djava.rmi.server.hostname=localhost -Dcom.sun.management.jmxremote.local.only=false"

GC_OPTS="-XX:+UseG1GC -XX:+UseStringDeduplication -XX:MaxGCPauseMillis=2000"

/bin/sh "$1"
