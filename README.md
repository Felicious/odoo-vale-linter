# Prose linter

## Install

### Windows
1. When opening terminal, select "Run as admin". The top of the terminal window should say: "Administrator: Command Prompt"
2. In terminal, run the [following command](https://docs.chocolatey.org/en-us/choco/setup#install-with-cmd.exe):
   ```shell
   @"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
   ```
3. Wait for it to download...
![installing choco](/screenshots/choco-install.png)
4. Then, type `choco` to confirm proper installation. You should see:
![successful install](/screenshots/choco-success.png)
5. Next, install vale by typing:
   ```shell
   choco install vale
   ```
   Type `y` and click `Enter` to accept the download
![vale install success](/screenshots/vale-success.png)
6. Navigate to the folder in terminal you'd like to download this [project](https://github.com/Felicious/odoo-vale-linter), and run
   ```shell
   git clone git@github.com:Felicious/odoo-vale-linter.git
   ```
7. Navigate into the new 'odoo vale linter' folder that was just created in terminal. Test our odoo proofreader works properly by typing `vale`

### Linux
1. Open terminal and `cd` to the folder you intend to store the vale project.
2. Download the [project](https://github.com/Felicious/odoo-vale-linter) by typing:
   ```shell
   git clone git@github.com:Felicious/odoo-vale-linter.git
   ```
3. Next, run a script to automatically download Vale, and type your password for your laptop when prompted.
   ```shell
   ./install.sh
   ```
   ![linux successful install](/screenshots/linux-install.png)
4. Once you're done, check Vale is correctly installed by typing:
   ```shell
   vale
   ```
![vale prompt](/screenshots/vale-success-msg.png)

## Add as VS code extension
1. Open VS code and select the square icon to search for extensions.
![open extension](/screenshots/extensions.png)
2. Search for an install the "Vale VCode" extension.
3. After installing, it's time to configure the vale extension to use our Odoo proofreading configurations. Open the VS code settings by clicking the gear icon > clicking Settings, or simply click "Ctrl + ,"
![open settings](/screenshots/open-settings.png)
4. In the User settings, select "Vale".
5. In the "Vale CLI: Config" field, enter the absolute path to the `vale.ini` file from the odoo-vale-linter project. To get that path, open the odoo-vale-linter project in VS code, right click the `vale.ini` file and select "Copy Path".
![copy path](/screenshots/vale-path.png)
6. In the "Vale CLI: Path" field, enter the path to the vale app installed using choco or homebrew. It's *probably* `/bin/vale`
![file configs](/screenshots/file-configs.png)

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