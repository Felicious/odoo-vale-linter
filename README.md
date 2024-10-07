# Odoo Documentation Vale Prose Linter

This project contains the [Vale](https://vale.sh/) configuration for linting prose on RST files in the Odoo documentation.

## 1. Clone This Repo

First, clone this project to the folder containing your GitHub repos:

```
github/
├── odoo/
│   └── documentation
└── odoo-vale-linter <-- this repo
```

Use the following command:

```shell
git clone https://github.com/Felicious/odoo-vale-linter.git
```

## 2. Install and Update Vale

Follow the installation instructions for your system here: https://vale.sh/docs/vale-cli/installation/#package-managers

### Update Vale

This repository uses Vale `3.7.0`. Please make sure this version, or a version above.

See the installed Vale version by running `vale --version` in the terminal.

To update Vale, run the following command in the terminal depending on your operating system:

- macOS: `brew upgrade vale`
- Linux: `sudo snap refresh vale`

## 3. Install Vale VS Code extension

Open VS Code and install the [Vale VSCode](https://marketplace.visualstudio.com/items?itemName=ChrisChinchilla.vale-vscode) extension.

Once installed, configure the vale extension to use the Odoo style configurations from this repository.

To do so, open the VS Code **Settings** and search for `Vale`.

In the **User** tab for **Vale**, select the **Enable Spellcheck** option in the **Vale Settings**, as well as set the **Readability Problem Location** field to **Both**.

Next, find the **Vale CLI: Config** field and enter the absolute path to the `.vale.ini` file from the **odoo-vale-linter** project.

![file configs](/screenshots/vscode-vale-config.png)

> ℹ️ Tip:</br>
> To get the absolute path of the `.vale.ini` file, open the **odoo-vale-linter** project in VS Code, right click the `.vale.ini` file and select **Copy Path** (do not click **Copy Relative Path**).</br></br>
> ![copy path](/screenshots/vale-path.png)

Next, enter the absolute path of Vale in the **Vale CLI: Path** field.

> ℹ️ Tip:</br>
> To get the absolute path of vale, type `which vale` in terminal and copy the result.</br></br>
> ![which vale](/screenshots/which-vale.png)

With these settings, the **Vale VSCode** extension should now be configured and [ready to use](#to-use) with the **odoo-vale-linter** project to lint prose on RST files.

## Use Vale in VS Code

When an RST file is saved, the alerts that Vale detects display both inline in the RST and in the **Problems** tab in VS Code.

> ℹ️ Tip:</br>
> Enable the **Problems** tab from **View --> Problems**.

Alert level severity:

1. :information_source: `Suggestion`: consider revising; double-check before ignoring.
2. :warning: `Warning`: requires attention; should not be ignored.
3. :x: `Error`:  critical format that must be addressed.


Either correct the RST or use discretion to ignore the suggestion.

![vale problems](/screenshots/vale-problems.png)

---

## To contribute

If a rule isn't behaving as expected, report it using the "Bug report" template under **Issues** --> **New issue**.

Alternatively, improve existing rules in the `styles/Odoo` folder, or create new ones. Don't forget to [include tests](./TESTING.md) for your rule in the `testdata` folder.

## Tested rules

1. **Spelling**: all words are spelled correctly
   - Odoo terms that shouldn't be flagged by spellcheck are found in the `styles/Vocab/Docs/accept.txt` folder
2. **Repeats**: flag any double words one after another
3. **Alt Tag**: Alt tag: flags that alt text should have a period
4. **Future Tense**: suggest to not use it
5. **Headings**: sentence case except for proper nouns and brands

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

## Blocks and limitations

1. To understand rst, Vale converts rst to html, then uses html to [check scopes](https://vale.sh/docs/topics/scoping/#types-formats-and-scopes). This means, to check whether all elements in a `guilabel`, an rst-specific tag, for example, we first have to build the doc (run `make fast` in terminal), *inspect* the guilabel element, and see what is output. Answer: :guilabel:`Confirm` is rendered as `<span class="guilabel">Confirm</span>`. Since there are no ways to define span classes from the [scope list](https://vale.sh/docs/topics/scoping/#markup), there isn't an easy way to target guilabel, menuselection, and other rst-specific tags ): This might be better done using the Python linter and adding to Sam's make review project instead.
2. To find out the part of speech of a sentence to set up a [sequence extension point](https://vale.sh/docs/topics/styles/#sequence) and control for grammar, run `vale tag [file name] en ""`. The endpoint (en "") is required to use this check.
3. `:abbr:` rendered in markup is an abbr HTML element, which is unsupported by Vale. I think all text in abbr are ignored, as I couldn't find a token
