import radius as radius

param rg string = resourceGroup().name

param sub string = subscription().subscriptionId

param registry string 

param version string

param magpieimage string 

resource env 'Applications.Core/environments@2022-03-15-privatepreview' = {
  name: 'corerp-resources-env-recipe-parameters-env'
  location: 'global'
  properties: {
    compute: {
      kind: 'kubernetes'
      resourceId: 'self'
      namespace: 'corerp-resources-env-recipe-parameters-env'
    }
    providers: {
      azure: {
        scope: '/subscriptions/${sub}/resourceGroups/${rg}'
      }
    }
    recipes: {
      'Applications.Link/mongoDatabases' :{
        mongodb: {
          templateKind: 'bicep'
          templatePath: '${registry}/test/functional/shared/recipes/mongodb-recipe-parameter:${version}'
          parameters: {
            documentdbName: 'acnt-operator-o-${rg}'
            mongodbName: 'mdb-operator-o-${rg}'
          }
        }
      }
    }
  }
}

resource app 'Applications.Core/applications@2022-03-15-privatepreview' = {
  name: 'corerp-rs-mongodb-recipe-parameters-old'
  location: 'global'
  properties: {
    environment: env.id
    extensions: [
      {
          kind: 'kubernetesNamespace'
          namespace: 'corerp-rs-mongodb-recipe-param-app-old'
      }
    ]
  }
}

resource webapp 'Applications.Core/containers@2022-03-15-privatepreview' = {
  name: 'mdb-param-ctnr-old'
  location: 'global'
  properties: {
    application: app.id
    connections: {
      mongodb: {
        source: recipedb.id
      }
    }
    container: {
      image: magpieimage
      env: {
        DBCONNECTION: recipedb.connectionString()
      }
      readinessProbe:{
        kind:'httpGet'
        containerPort:3000
        path: '/healthz'
      }
    }
  }
}

resource recipedb 'Applications.Link/mongoDatabases@2022-03-15-privatepreview' = {
  name: 'mdb-recipe-param-db-old'
  location: 'global'
  properties: {
    application: app.id
    environment: env.id
    recipe: {
      name: 'mongodb'
      parameters: {
        documentdbName: 'acnt-developer-o-${rg}'
      }
    }
  }
}
