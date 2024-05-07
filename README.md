# Odoo Documentation Vale Prose Linter

This project contains the [Vale](https://vale.sh/) configuration for linting prose on RST files in the Odoo documentation.

## Install

The following Mac, Windows, and Linux sections contain instructions on how to install Vale and configure it to use the **odoo-vale-linter** project as a style guide.

### Mac

To configure Vale on Mac, [Homebrew](https://brew.sh/) needs to be installed. To install Homebrew, run the following command in the terminal:

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Once Homebrew is installed, run the following command in the terminal to [install Vale](https://vale.sh/docs/vale-cli/installation/#package-managers):

```shell
brew install vale
```

To check that Vale is installed, run the following command in the terminal:

```shell
vale --version
```

Next, in the terminal, navigate to your `/github` folder using the `cd <path>` command to change your active directory. Once in `/github`, clone this repository to your local machine with the following command:

```shell
git clone https://github.com/Felicious/odoo-vale-linter.git
```

With the repository cloned, proceed to the [Add as VS Code extension](#add-as-vs-code-extension) section.

### Windows

To configure Vale on Windows, [Chocolatey](https://chocolatey.org/) needs to be installed. To install Chocolatey, open a new terminal window as an administrator.

> ℹ️ Tip:</br>
> When opening terminal, select **Run as administrator**. The top of the terminal window should say: "Administrator: 'shell'".

The installation commands depend on the command-line shell being used by your terminal:

- [Install with cmd.exe](https://docs.chocolatey.org/en-us/choco/setup#install-with-cmd.exe)
- [Install with PowerShell.exe](https://docs.chocolatey.org/en-us/choco/setup#install-with-powershell.exe)

Once the installation script is ran, wait for Chocolatey to download.

![installing choco](/screenshots/choco-install.png)

Then, type `choco` to confirm proper installation. The installed version should be retured.

![successful install](/screenshots/choco-success.png)

Next, install Vale by typing:

```shell
choco install vale
```

Type `y` and press `Enter` to accept the download.

![vale install success](/screenshots/vale-success.png)

To check that Vale is installed, run the following command in the terminal:

```shell
vale --version
```

Next, navigate to your `/github` folder using the `cd <path>` command to change your active directory. Once in `/github`, clone this repository to your local machine with the following command:

```shell
git clone https://github.com/Felicious/odoo-vale-linter.git
```

With the repository cloned, proceed to the [Add as VS Code extension](#add-as-vs-code-extension) section.

### Linux
1. Open terminal and go to your `/github` folder using the `cd <path>` command to change your active directory.
2. Once in `/github`, clone this [repository](https://github.com/Felicious/odoo-vale-linter) to your local machine by typing:
   ```shell
   git clone git@github.com:Felicious/odoo-vale-linter.git
   ```
3. Next, `cd` to `github/odoo-vale-linter`, run the following code, and enter your password when prompted:
   ```shell
   ./install.sh
   ```
   ![linux successful install](/screenshots/linux-install.png)
> ℹ️ Tip:</br>
> When typing your password, the terminal doesn't show anything you type, including ***.

4. Once you're done, check Vale is correctly installed by typing:
   ```shell
   vale --version
   ```
![vale prompt](/screenshots/vale-success-msg.png)

It should display "version 3.0.5".

## Update

This repository uses Vale `3.0.5`. Please make sure this version, or a version above, of Vale is installed by running `vale --version` in the terminal.

To update Vale, run the following command in the terminal depending on your operating system:

- Mac: `brew upgrade vale`
- Windows: `choco upgrade vale`
- Linux: `./install.sh`

To update to the latest changes from this repository, simply pull the latest changes from the `master` branch by navigating to the **odoo-vale-linter** project folder in the terminal, and running the following command:

```shell
git pull origin master
```

## Add as VS Code extension

Open Visual Studio Code and install the [Vale VSCode](https://marketplace.visualstudio.com/items?itemName=ChrisChinchilla.vale-vscode) extension.

Once installed, it's time to configure the vale extension to use the Odoo style configurations from the **odoo-vale-linter** repository.

To do so, open the VS Code **Settings** and search for `Vale`. In the **User** tab for **Vale**, find the **Vale CLI: Config** field and enter the absolute path to the `.vale.ini` file from the **odoo-vale-linter** project.

![file configs](/screenshots/file-configs.png)

> ℹ️ Tip:</br>
> To get the absolute path of the `.vale.ini` file, open the **odoo-vale-linter** project in VS Code, right click the `.vale.ini` file and select **Copy Path**.</br>
> ![copy path](/screenshots/vale-path.png)

Be sure to also select the **Enable Spellcheck** option in the **Vale Settings**, as well as set the **Readability Problem Location** field to **Both**.

Then, enter the absolute path of Vale in the **Vale CLI: Path** field.

> ℹ️ Tip:</br>
> To get the absolute path of vale on a Linux or Mac OS, type `which vale` in terminal and copy the result.</br>
> ![which vale](/screenshots/which-vale.png)

With these settings, the **Vale VSCode** extension should now be configured and ready [to use](#to-use) with the **odoo-vale-linter** project to lint prose on RST files.

## To use
### VS code extension
1. Save a file (Ctrl + s) to make the suggestions show up.
2. The number that appears in the tab of the file in VS code is the number of problems found.
![num of issues](/screenshots/problems.png)
3. Click the bar at the bottom to open the tab to view all of the things the tool found.
![alerts](/screenshots/alerts.png)
4. In the pop-up that opens, "problems" with the yellow alert symbols are warnings and things the writer should probably change. The blue info icon are suggestions for brevity and word choice that are up to the writer's discretion whether to apply the changes or not. Red 'x' icons represent errors that **must** be fixed.
![how-to-use](/screenshots/problems-found.png)

### Terminal version
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

## Blocks and limitations ):
1. To understand rst, Vale converts rst to html, then uses html to [check scopes](https://vale.sh/docs/topics/scoping/#types-formats-and-scopes). This means, to check whether all elements in a `guilabel`, an rst-specific tag, for example, we first have to build the doc (run `make fast` in terminal), *inspect* the guilabel element, and see what is output. Answer: :guilabel:`Confirm` is rendered as `<span class="guilabel">Confirm</span>`. Since there are no ways to define span classes from the [scope list](https://vale.sh/docs/topics/scoping/#markup), there isn't an easy way to target guilabel, menuselection, and other rst-specific tags ): This might be better done using the Python linter and adding to Sam's make review project instead.
2. To find out the part of speech of a sentence to set up a [sequence extension point](https://vale.sh/docs/topics/styles/#sequence) and control for grammar, run `vale tag [file name] en ""`. The endpoint (en "") is required to use this check.
3. `:abbr:` rendered in markup is an abbr HTML element, which is unsupported by Vale. I think all text in abbr are ignored, as I couldn't find a token
