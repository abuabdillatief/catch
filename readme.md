# catch
__*catch*__ is a simple logging package that colorizes logs to help you read errors more easily and less frustrating.


# Install:

    go get github.com/abuabdillatief/catch

# Examples:

    type PrintType string

    const (
        TypeError PrintType = "Error"
        TypeWarn  PrintType = "Warn"
        TypeInfo  PrintType = "Info"
    )
## Simple logging
  	err := errors.New("mongo: no documents in result")
	catch.Error(err, "cant find document")
	catch.Warn(err, "level 1 warning")
	catch.Inform(err)

![example of simple logigng](./assets/image.png)

## Custom logging
    var customLog = make(map[string]string)
    customLog["count"] = "first try"
    customLog["heads"] = "eeve"

    catch.CustomLog(customLog, catch.TypeError)
![example of custom loging](./assets/custom_log.png)

## Create, save and delete log file
    err := errors.New("mongo: no documents in result")
	c := catch.NewLog("catch")
	c.SaveToLogFile(err)
After saving log file, file will look like this:

![example of custom loging](./assets/log_file.png)

To delete log file, simply call:

    c.DeleteLogFile()

# Credits
- [Color by fatih](https://github.com/fatih/color)


