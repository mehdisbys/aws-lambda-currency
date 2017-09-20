# AWS Lambda currency API fetch with Golang


Running Golang on AWS using the shim provided by [Apex](https://github.com/apex/apex).

The code structure respects the [Apex folder structure](https://github.com/apex/apex#example)

### Purpose

I have written this code to try AWS Lambda capabilities using Golang. This code goes and fetch the current currency values with USD dollars as the reference.

### Requirements

- Have an account on AWS 
- Setup an IAM role
- Configure the credentials on your local machine
- Get a free API key from [currencylayer.com](http://currencylayer.com) and put in an .env with key CURRENCY_API_KEY

### Try it

Once the requirements above are fulfilled, run the following command :

`apex invoke golangHello`

or use the following PHP script (coming soon)
