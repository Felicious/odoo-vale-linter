# Prose linter

## Install
1. [Install Vale](https://vale.sh/docs/vale-cli/installation/).
2. Fork this repository.

## To use
1. In terminal, navigate to the folder where the `.vale.ini` file is located.
2. Copy the absolute file of the file you'd like to run the test on. To do that, right-click   the
   desired RST file in VSCode, and select the **Copy Path** option from the drop-down menu.

3. In terminal, type `vale [absolute file path]`

## Working rules
1. **Alt Tag**: Raises an error when there isn't a period at the end of an image's alt tag.
2. **No Future Case**: Suggests to not include "will" in sentences, unless it's to suggest potential options.

   <details>
   <summary>Allowed example</summary>
   ``Selecting the :guilabel:`Todo` filter option will generate a list of tasks to-do, with ...``
   </details>

3. **Sentence Case**: Enforces that all headings only have the first letter capitalized. Exceptions: proper nouns (WIP)
4. **Passive Voice**: Reminder to remove uncertain words like "may, possible, could" and keep the writing focused on actively instructing the reader to do something.

WIP rules
1. Enforce sentence case on headings. Exceptions: capitalize Odoo if it's not the first word. Refer to this [doc](https://vale.sh/docs/topics/styles/#capitalization)
   - Exception: Odoo and abbreviations
2. Raise an alert when an acronym hasn't been defined yet. Refer to this [doc](https://vale.sh/docs/topics/styles/#conditional)
3. Product formatting: capitalize and italicize app names
4. Punctuation>quotation marks: quotation marks-- chances are we shouldn't see them (warning). Exception: exercise block
5. If the word "click" directly precedes bold or italic text then should receive gui warning
6. Error: hyperbole
  utmost, paramount, etc (can find you more)
7. Error: sales jargon
  TBD

## Blocks and limitations ):
1. To understand rst, Vale converts rst to html, then uses html to [check scopes](https://vale.sh/docs/topics/scoping/#types-formats-and-scopes). This means, to check whether all elements in a `guilabel`, an rst-specific tag, for example, we first have to build the doc (run `make fast` in terminal), *inspect* the guilabel element, and see what is output. Answer: :guilabel:`Confirm` is rendered as `<span class="guilabel">Confirm</span>`