resource app 'radius.dev/Application@v1alpha3' = {
  name: 'kubernetes-resources-mongo'
  
  resource webapp 'ContainerComponent' = {
    name: 'todomongo'
    properties: {
      container: {
        image: 'radius.azurecr.io/magpie:latest'

        // This is here so we make sure to exercise the 'programmatic secrets' code path.
        env: {
          DB_CONNECTION: mongodb.connectionString()
        }
      }
      connections: {
        mongodb: {
          kind: 'mongo.com/MongoDB'
          source: mongodb.id
        }
      }
    }
  }

  resource mongodb 'mongodb.com.MongoDBComponent' = {
    name: 'mongodb'
    properties: {
        managed: true
    }
  }
}
