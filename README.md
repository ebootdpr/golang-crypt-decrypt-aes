# golang-crypt-decrypt-aes
this ./cry files encryps everything except
- /node_modules (folder)
- /.git (folder) (it renames it to enc.git to avoid pushing decrypted files)
- /public (folder)
- ./cry-name (file executable name)

Only tested on ubuntu.

## Usage
Put the cry file on the folder you want to encrypt
execute the cry file, 
- insert your alphanumerik key, 32 max (not tested with nonASCII).
- insert your secret vectors numerics from 0 to 9, 16 max.
it creates a file to check if the folder is already encripted, dont delete
