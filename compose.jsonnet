{
  version: '3.1',

  local user = 'garmin',
  local password = 'garmin',
  services: {
    db: {
      image: 'docker.io/library/postgres',
      restart: 'always',

      environment: {
        POSTGRES_DB: user,
        POSTGRES_PASSWORD: password,
        POSTGRES_USER: user,
      },
      ports: ['${POSTGRES_PORT:-5432}:5432'],
      volumes: ['/home/jdb/garmin:/var/lib/postgresql/data:z'],
    },
    grafana: {
      image: 'docker.io/grafana/grafana',
      restart: 'always',

      entrypoint:
        local dataSources = {
          apiVersion: 1,
          datasources: [{
            name: 'garmin',

            database: user,
            url: 'db',  // host
            jsonData: {
              sslmode: 'disable',
            },
            isDefault: true,
            secureJsonData: {
              password: password,
            },
            type: 'postgres',
            user: user,
          }],
        };
        [
          'sh',
          '-euc',
          |||
            printf "%s" > /etc/grafana/provisioning/datasources/garmin.yml
            exec /run.sh
          ||| % std.manifestYamlDoc(dataSources),
        ],
      depends_on: ['db'],
      environment: [
        'GF_AUTH_ANONYMOUS_ENABLED=true',
        'GF_AUTH_ANONYMOUS_ORG_ROLE=Admin',
        'GF_DATABASE_TYPE=postgres',
        'GF_DATABASE_HOST=db',
        'GF_DATABASE_USER=' + user,
        'GF_DATABASE_PASSWORD=' + password,
        'GF_DATABASE_SSL_MODE=disable',
        'GF_INSTALL_PLUGINS=pr0ps-trackmap-panel',
        'GF_PLUGINS_PLUGIN_ADMIN_ENABLED=true',
      ],
      ports: ['${GRAFANA_PORT:-3000}:3000'],
    },
  },
}
