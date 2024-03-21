# README

An app I shitcoded while trying to play around with go and svelte and checking software dev frameworks. 

## How it works

Goal: To be able to swiftly find and send images to chats

1. Select a folder (default is home)
2. Scans all the images
3. Creates meme-folder.json in the folder with each image maped

- add and remove from the interface tags for each emage
- search images by tags now
- use it with cloud to make it available on multiple devices
- add linked images and copy their links when needed

## 0.2
- Functionality: add an image by link in the meme-folder.json file - done

## TODO features
- Fix for windows - oh hold until wails 3
- make .deb to autoinstall dependencies on linux
- Option to upload images on imgur and return the link (after upload transform the upload button )
- Option to download only link available images (after download transform the download button in open in file)
- Option to simulate copy-paste of the images and drag and drop in desired chat

## TODO code
- Better naming
- Optimize go code with pointers

## About

This is the official Wails Svelte template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

