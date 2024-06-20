# Ascii Art Web Export

## Project Description
Ascii Art Web Export is a web application that allows users to create ASCII art and export it as a downloadable file. The application converts user input text into ASCII art using three different formats (thinkertoy, shadow, or standard) and provides the output as a downloadable file.

## Features
- Create ASCII art in three formats: thinkertoy, shadow, or standard.
- Download the generated ASCII art as a `.txt` file.
- Ensure correct permissions (read and write) for the downloaded file.
- Use HTTP headers (Content-Type, Content-Length, Content-Disposition) for file transfer over the web.

## Requirements
- Go 1.16 or later
- A web browser

## Installation
1. Clone the project:
    ```sh
    git clone https://github.com/username/project.git
    cd project
    ```

2. Prepare the necessary files (thinkertoy.txt, shadow.txt, standard.txt):
    - These files should contain ASCII art characters.

3. Install Go dependencies and run the project:
    ```sh
    go mod init project
    go mod tidy
    go run main.go
    ```

4. Open your web browser and go to `http://localhost:8080`.

## Usage
1. Enter the text you want to convert to ASCII art in the text box.
2. Select the banner type (thinkertoy, shadow, or standard).
3. Check the `Download` box to download the ASCII art and click the `Submit` button.
4. The ASCII art will be downloaded as a file with the name `ascii_art.txt`.

## Project Files
- `main.go`: The main code file of the application.
- `form.html`: HTML form file for the user interface.
- `style.css`: Style file for site design.
- `thinkertoy.txt`, `shadow.txt`, `standard.txt`: ASCII art character files.

## Example Usage
### Web Interface
1. Go to `http://localhost:8080` in your web browser.
2. Enter `Hello, World!` in the text box.
3. Select `standard` as the banner type.
4. Check the `Download` box and click `Submit`.
5. The `ascii_art.txt` file will be downloaded, containing the ASCII art for `Hello, World!`.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


## Contact
For any questions or feedback, please contact:
- **Email**: ahatdemirezenn@gmail.com
- **GitHub**: [Ahat Demirezen](https://github.com/ahatdemirezen)
