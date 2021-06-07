# catch
__*catch*__ is a simple logging package that colorizes logs to help you read errors more easily and less frustrating.


# Install:

    go get github.com/abuabdillatief/catch

# Examples:

### Simple logging
  	err := errors.New("mongo: no documents in result")
	catch.Error(err, "cant find document")
	catch.Warn(err, "level 1 warning")
	catch.Inform(err)

![example of simple logigng](./assets/image.png)

### Custom logging
    var customLog = make(map[string]string)
    customLog["count"] = "first try"
    customLog["heads"] = "eeve"

    catch.CustomLog(customLog, catch.TypeError)
![example of custom loging](./assets/custom_log.png)

# Credits
- [Color by fatih](https://github.com/fatih/color)


