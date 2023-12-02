# Prose linter

## To use
1. Copy the absolute file of the file you'd like to run the test on. To do that, right-click   the
   desired RST file in VSCode, and select the **Copy Path** option from the drop-down menu.

2. In terminal, type `vale [absolute file path]`

## Working rules
1. **PeriodAlt**: Raises an error when there isn't a period at the end of an image's alt tag.

WIP rules
1. Enforce sentence case on headings. Exceptions: capitalize Odoo if it's not the first word. Refer to this [doc](https://vale.sh/docs/topics/styles/#capitalization)
2. Raise an alert when an acronym hasn't been defined yet. Refer to this [doc](https://vale.sh/docs/topics/styles/#conditional)
