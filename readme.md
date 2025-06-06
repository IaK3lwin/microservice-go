# Microservice to lear golang

## Requerimenrs
```sh

- () the api should be able to create a new category
- () the api should be able to list all category
- () the api should be able to get a category
- () the api should be able to delete a category
- () the api should be able to update a existent  categoy
- () the api should send a event to a broker when a category name is update
```

## Next steps
```sh

- () dockerize
- () Permission roles
- () deploy

```

## folder pattener in golang community
```sh
	./
	|
	|--- **cmd**
    |------- api
	|--------- main.go
    |--- internal -> all files related to tour application's file stay here | this is all about community standards
    |--------- entities -> all entities in the project
    |------------ category.go


```

## FramWorks used
### - Gin (http router with our routes )
```sh

    ``` 
    ...

        router := gin.Default() //Default is a instance in the gin's router with middleware defualt and preconfigured.

        router.GET("/<route>", HandleFunc) //.get() is a verb in the http protocol that get a route and a handle function, which is called wh when tha route is called

     ```




```

### - Air-Verse Live reload for Go app

    command to install: `go install github.com/air-verse/air@latest`

#### start air
    for init air use comand: "air init"
#### config  
    a file called .air.toml is created
    This is where you will point to the boot file

    ```
    cmd = "go build -o .temp/main <path do main.go> "

    ```

#### init air
    To ru the server with Air usethe command: `air`


## Folder Entities

---
    title: Entities
    description: "stay  entities from project"


---

### Category

    the  category structure  has some fields like:
        Id uint
        Name string
        CreateData time.Time
        UpdateData time.Time

    #### The community stardard is to  create a function to instantiate  category structure

    ```
        // create category recommended
func (categ Category) CreateCategory(Nome string) (*Category, error) {
	category := Category{
		Nome:        Nome,
		DataCreated: time.Now(),
		DataUpdate:  time.Now(),
	}

	//validar
	if err := category.isValid(); err != nil {
		//retorna um erro
		return nil, err
	}

	return &category, nil
}
    ```
