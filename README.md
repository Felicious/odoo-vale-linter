# Prose linter

## To use
1. Copy the absolute file of the file you'd like to run the test on. To do that, right-click   the
   desired RST file in VSCode, and select the **Copy Path** option from the drop-down menu.

2. In terminal, type `vale [absolute file path]`

## Working rules
1. **Alt Tag**: Raises an error when there isn't a period at the end of an image's alt tag.
2. **No Future Case**: Suggests to not include "will" in sentences, unless it's to suggest potential options.

   <details>
   <summary>Allowed example</summary>
   ``Selecting the :guilabel:`Todo` filter option will generate a list of tasks to-do, with ...``
   </details>



WIP rules
1. Enforce sentence case on headings. Exceptions: capitalize Odoo if it's not the first word. Refer to this [doc](https://vale.sh/docs/topics/styles/#capitalization)
   - Exception: Odoo and abbreviations
2. Raise an alert when an acronym hasn't been defined yet. Refer to this [doc](https://vale.sh/docs/topics/styles/#conditional)
3. Product formatting: capitalize and italicize app names
4. Punctuation>quotation marks: quotation marks-- chances are we shouldn't see them (warning)
5. If the word "click" directly precedes bold or italic text then should receive gui warning
6. Passive voice: may, would, could, should, might, possible
  "Transferring stock from one warehouse to another is possible in Odoo"
  "stick to active instruction"
7. Error: hyperbole
  utmost, paramount, etc (can find you more)
8. Error: sales jargon
  TBD
9. Capitalize everything in a guilabel


