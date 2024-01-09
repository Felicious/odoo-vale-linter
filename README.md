# Prose linter

## Install

1. [Install Vale](https://vale.sh/docs/vale-cli/installation/).
2. Download this project
   ```shell
   git clone git@github.com:Felicious/odoo-vale-linter.git
   ```

## To use

1. In terminal, navigate to the folder where the `.vale.ini` file is located.
2. Copy the absolute file of the file you'd like to run the test on. To do that, right-click the
   desired RST file in VSCode, and select the **Copy Path** option from the drop-down menu.

3. In terminal, type `vale [absolute file path]`

## To contribute
1. If a rule isn't behaving as expected, report it using the "Bug report" template under **Issues** --> **New issue**.
2. Alternatively, improve existing rules in the `styles/Odoo` folder, or create new ones.
3. Don't forget to [include tests](./TESTING.md) for your rule in the `testdata` folder.

## Tested rules
1. **Spelling**: all words are spelled correctly
   - Odoo terms that shouldn't be flagged by spellcheck are found in the `styles/Vocab/Docs/accept.txt` folder
2. **Repeats**: flag any double words one after another
3. **Alt Tag**: Alt tag: flags that alt text should have a period
4. **Future Tense**: suggest to not use it

WIP rules
1. Enforce sentence case on headings. Exceptions: capitalize Odoo if it's not the first word. Refer
   to this [doc](https://vale.sh/docs/topics/styles/#capitalization)
   - Exception: Odoo and abbreviations
2. Raise an alert when an acronym hasn't been defined yet. Refer to this
   [doc](https://vale.sh/docs/topics/styles/#conditional)
3. Product formatting: capitalize and italicize app names
4. Punctuation>quotation marks: quotation marks-- chances are we shouldn't see them (warning).
   Exception: exercise block
5. If the word "click" directly precedes bold or italic text then should receive gui warning
6. Error: hyperbole
  utmost, paramount, etc (can find you more)
7. Error: sales jargon: TBD

## Blocks and limitations ):
1. To understand rst, Vale converts rst to html, then uses html to [check scopes](https://vale.sh/docs/topics/scoping/#types-formats-and-scopes). This means, to check whether all elements in a `guilabel`, an rst-specific tag, for example, we first have to build the doc (run `make fast` in terminal), *inspect* the guilabel element, and see what is output. Answer: :guilabel:`Confirm` is rendered as `<span class="guilabel">Confirm</span>`. Since there are no ways to define span classes from the [scope list](https://vale.sh/docs/topics/scoping/#markup), there isn't an easy way to target guilabel, menuselection, and other rst-specific tags ): This might be better done using the Python linter and adding to Sam's make review project instead.
2. To find out the part of speech of a sentence to set up a [sequence extension point](https://vale.sh/docs/topics/styles/#sequence) and control for grammar, run `vale tag [file name] en ""`. The endpoint (en "") is required to use this check.
3. `:abbr:` rendered in markup is an abbr HTML element, which is unsupported by Vale. I think all text in abbr are ignored, as I couldn't find a token