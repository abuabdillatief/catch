# catch
__*catch*__ is a simple logging package that colorizes logs to help you read errors more easily and less frustrating.


# Install:

    go get github.com/abuabdillatief/catch

# Examples:

  	err := errors.New("mongo: no documents in result")
	catch.Error(err, "cant find document")
	catch.Warn(err, "level 1 warning")
	catch.Inform(err)

![example of catch.Error](./assets/image.png)

# Credits
- [Color by fatih](https://github.com/fatih/color)

