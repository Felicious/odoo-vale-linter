# Odoo Documentation Vale Prose Linter

This project contains the [Vale](https://vale.sh/) configuration for linting prose on RST files in the Odoo documentation.

## Alert Levels

- 🔵 **Suggestion**: Double check this– but not required to fix.
- 🟡 **Warning**: This should probably be fixed.
- 🔴 **Error**: This should absolutely be fixed.

## Odoo Styles

- 🔴 [Alt Tag](./styles/Odoo/AltTag.yml): Alt tags should end in a period.
- 🟡 [App Names](./styles/Odoo/AppNames.yml): Odoo app names should use bold formatting.
- 🟡 [Avoid](./styles/Odoo/Avoid.yml): Don't use phrases like "obviously, simply, quickly".
- 🟡 [Common Nouns](./styles/Odoo/CommonNouns.yml): Don't capitalize common nouns (unless directly referring to the UI, or part of active instructions).
- 🟡 [Contraction](./styles/Odoo/Contraction.yml): Avoid using contractions whenever possible.
- 🟡 [Future Tense](./styles/Odoo/FutureTense.yml): Avoid writing in future tense.
- 🔵 [Icons](./styles/Odoo/Icons.yml): The preferred descriptor for each icon. Always use the tooltip description if present.
- 🟡 [Image Align](./styles/Odoo/ImageAlign.yml): Consider removing `:align:` attribute from images.
- 🔴 [Nonwords](./styles/Odoo/Nonwords.yml): Common typos of English words.
- 🟡 [Optional Plurals](./styles/Odoo/OptionalPlurals.yml): Avoid writing a sentence that adds an optional plural (s) to words.
- 🟡 [Ordinals](./styles/Odoo/Ordinals.yml): Write out ordinal numbers.
- 🟡 [Pronouns](./styles/Odoo/Pronouns.yml): Avoid using pronouns.
- 🟡 [SmartQuotes](./styles/Odoo/SmartQuotes.yml): Avoid using smartquotes (e.g. `“ ”` or `‘ ’`).
- 🔴 [Spacing](./styles/Odoo/Spacing.yml): Don't use double spaces.
- 🔴 [Spelling](./styles/Odoo/Spelling.yml): Common words that are not part of the default dictionary.
- 🟡 [Start Sentence](./styles/Odoo/StartSentence.yml): Avoid starting a paragraph with "But", "There is", or "There are".
- 🔴 [WordList](./styles/Odoo/WordList.yml): Use the preferred format for these words.

## To contribute

If a rule isn't behaving as expected, report it using the "Bug report" template under [Issues](https://github.com/Felicious/odoo-vale-linter/issues).

Alternatively, improve existing rules in the `styles/Odoo` folder, or create new ones. Don't forget to [include tests](./TESTING.md) for your rule in the `testdata` folder.


# Install Steps

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

## 2. Install Vale

Follow the installation instructions for your system:

- **macOS**: `brew install vale`
- **Linux**:

  1. In a **Terminal**, navigate to the `odoo-vale-linter` repository using: `cd path/to/repo`.
  2. Run the `./install.sh` command, and enter your password when prompted.

This repository uses Vale `3.9.0`. Please make sure this version, or a version above.

See the installed Vale version by running `vale --version` in the terminal.

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
> To get the absolute path of vale, type `which vale` in terminal and copy the result.
> <details>
> <summary>macOS</summary>
>
> ![which vale](/screenshots/which-vale-mac.png)
>
> </details>
> <details>
> <summary>Linux</summary>
>
> ![which vale](/screenshots/which-vale-linux.png)
>
> </details>

With these settings, the **Vale VSCode** extension should now be configured and [ready to use](#to-use) with the **odoo-vale-linter** project to lint prose on RST files.

## Update your local repo

Be sure to frequently update your local repo to the latest changes.

First, open a new **Terminal** window.

Then navigate to the directory containing this repo, for example:

```
cd github/odoo-vale-linter
```

Next, pull the latest changes from origin:

```
git pull origin main
```

After the pull, your local instance of Vale will enforce the most up-to-date styles.

> ℹ️ Tip:</br>
> To be notified of changes to this repo, click the <svg aria-hidden="true" focusable="false" class="octicon octicon-eye" viewBox="0 0 16 16" width="16" height="16" fill="currentColor" style="display: inline-block; user-select: none; vertical-align: text-bottom; overflow: visible;"><path d="M8 2c1.981 0 3.671.992 4.933 2.078 1.27 1.091 2.187 2.345 2.637 3.023a1.62 1.62 0 0 1 0 1.798c-.45.678-1.367 1.932-2.637 3.023C11.67 13.008 9.981 14 8 14c-1.981 0-3.671-.992-4.933-2.078C1.797 10.83.88 9.576.43 8.898a1.62 1.62 0 0 1 0-1.798c.45-.677 1.367-1.931 2.637-3.022C4.33 2.992 6.019 2 8 2ZM1.679 7.932a.12.12 0 0 0 0 .136c.411.622 1.241 1.75 2.366 2.717C5.176 11.758 6.527 12.5 8 12.5c1.473 0 2.825-.742 3.955-1.715 1.124-.967 1.954-2.096 2.366-2.717a.12.12 0 0 0 0-.136c-.412-.621-1.242-1.75-2.366-2.717C10.824 4.242 9.473 3.5 8 3.5c-1.473 0-2.825.742-3.955 1.715-1.124.967-1.954 2.096-2.366 2.717ZM8 10a2 2 0 1 1-.001-3.999A2 2 0 0 1 8 10Z"></path></svg> **Watch** button on the top-right of the repo, then select **All Activity**.

---

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
