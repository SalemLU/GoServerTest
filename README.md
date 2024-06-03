# Go Tips

- Use Visual Studio Code with the Go module installed so it can handle imports and formating for you on the fly, as well as show you details on available variable members and method details
- In the go.mod file, the first line (module) is the name your folder/package will take when other apps try to reference it
- The proto file will make the bulk of the backend for you, and if you modify it to include more methods or messages, re-running the protoc command will add those in. Do not modify the generated code
- '*' is used to declare a pointer and to dereference, and '&' points to the memory address of the stored value

<br />

# How to run the apps

Run the server app first with ```go run main.go struts.go```<br />
Then run the web app with ```go run main.go```. The site is accessible via http://localhost:8000. Look into the "db.json" file for valid films to try request for the director's name.
