# minds-go-sdk 

## Getting started
1. Initialize the client

    Intialize the client with your api key to interact with the sdk.

    ```go
    import "github.com/ashish111333/minds-go-sdk/client"

    
    func main(){
        client,err:=client.NewClient("YOUR_API_KEY)
        if err!=nil{
            fmt.println("failed to create client",err)
        }
    }
    ```
2. Creating a Datasource

    You can connect to various databases, such as PostgreSQL, by configuring your data source. Use the DatabaseConfig to define the connection details for your data source.
    
    ```go
    import "github.com/ashish111333/minds-go-sdk/datasources"

    func main(){
        postgres_config:=datasources.Datasource{
            DatabaseConfig:{
                   name='my_datasource',
                    description='<DESCRIPTION-OF-YOUR-DATA>',
                    engine='postgres',
                    connection_data={
                        'user': 'demo_user',
                        'password': 'demo_password',
                        'host': 'samples.mindsdb.com',
                        'port': 5432,
                        'database': 'demo',
                        'schema': 'demo_data'
                    },
                    tables=['<TABLE-1>', '<TABLE-2>']
                }
        }
        
    }
    ```
3. Creating a mind
   
   You can create a mind and associate it with a data source.
   










## Managing Minds

You can create a mind or replace an existing one with the same name.

```go
 ```

### List minds

get list of minds you created

```go
    
    minds,err:=client.minds.list()
    
```
### get mind by name

02:01 - 15m ahead 
you can fetch details of a mind by its name.

```go 
mind,err:=client.minds.get("mind_name")
```

### Remove mind
you can remove a mind given it's name

```go
err=client.minds.drop("mind_name")
````