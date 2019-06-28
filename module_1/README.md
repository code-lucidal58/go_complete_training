# Module 1

## Text templates 
Templates are a way to define a format with variables. Later the variables are substituted as per requirement. This is
very helpful code/script generation is to be done. They are going when format is to kept separated from data. It allows 
you to print values, select fields, call functions, do control flow with if-else and range. The package used is 
```text/template```

Surround variables wit {{}} and add a dot(.) before the variable name. The names should be in capital letters to be used
in template. ```template.Must()``` is a static function that validates the contents of the template. Panics if there are
any error. Validations like matching braces, etc. 

Some important sections from golang text/template package.
* [Text and spaces](https://golang.org/pkg/text/template/#hdr-Text_and_spaces)