![password-manager](resources/logo.png) 
# Password Manager

[![Build Status](https://travis-ci.com/ThilinaManamgoda/password-manager.svg?branch=master)](https://travis-ci.com/ThilinaManamgoda/password-manager)

Encrypt your passwords into a file and have access with ease.
### Synopsis

A local password manager which simply encrypts your passwords in to a file and searching passwords made easy ![how-to-install](resources/exclamation-mark.png).

![how-to-install](resources/instalation.png) 
### How to install

#### Ubuntu ![Ubuntu](resources/linux.png)

1. Using the command line, add the following to your /etc/apt/sources.list system config file,

    `sudo echo "deb https://dl.bintray.com/maanafunedu/maanadev-debian stable main" | sudo tee -a /etc/apt/sources.list`
1. Configure Bintray public GPG key,

    `sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61`
1. Update repository,   
      
     ```sudo apt update```
1. Install password-manager, 
 
    ```sudo install password-manager```

#### MacOS ![macos](resources/macos.png)

1. Tap into password-manager brew formula with the following command,
        
     ```brew tap ThilinaManamgoda/homebrew-password-manager```
1. Install password-manager, with the following command,  
       
     ```brew install password-manager```
     
![How to use](resources/how-to-use.png)   
### How to use

1. Initialize the password manager with the following command,
    ```$xslt
    password-manager init
    ```
    Enter the Master password once prompted.
    
2. Add a password entry with the following command,
    ```$xslt
    password-manager add TEST -i
    ```
    Here password entry id is **TEST** which always should be a unique value. By passing **-i** parameter we are enabling
    the **Interactive** mode. Once this command is executed, the following entries will be listed,
    
    ```$xslt
        Username: <Enter the username for this password entry>
        Password: <Enter the password for this password entry>
        Enter the Password again: <Enter the password again for this password entry>
        Lables: <Enter lables that can be used for searching. This can be list of comma seperated values>
        Master password: <Enter Master password>
        
        Example:
        
        Username: username@test.com
        Password: ***********
        Enter the Password again: ***********
        Lables: test,first,firstTime
        Master password: <Enter Master password>
    ```

1. Get a password entry with **ID**
    ```$xslt
        password-manager get <ID> 
    ```     
    Enter the Master password once prompted. For an example,
    ```$xslt
        password-manager get TEST
    ```
1. Search a password entry with a **ID**
    ```$xslt
        password-manager search <ID>
    ```
    Enter the Master password once prompted. For an example,
    ```$xslt
        password-manager search test
    ```
    Once you enter the Master password the a list of password entries will be listed that match the given label. Once the entry is selected, the **password** will be copied to the clipboard.
     
1. Search a password entry with a **label**
    ```$xslt
        password-manager search -l <LABEL>
    ```
    Enter the Master password once prompted. For an example,
    ```$xslt
        password-manager search -l test
    ```
    Once you enter the Master password the a list of password entries will be listed that match the given label. Once the entry is selected, the **password** will be copied to the clipboard.
 
### Configuration

1. Follow the instructions to override default configurations using a configuration file.

    1. Create a password-manager.yaml file with following content. Keep configurations that need to be overridden. 
        ```$xslt
       # Set encryptor for encryping passwords. Ex: "AES"
       encryptorID: "AES"
       # Maximum list size for selection drop down.
       selectListSize: 5
       # Directory which holds the Password manager related files.
       directoryPath: "~/password-manager"
       # Set storage configurations.
       storage:
         # File storage configuration.
         file:
           # Enable File storage.
           enable: true
           # Password Database file name.
           passwordDBFile: "password_db"
           # File permission for given file in the path.
           permission: 0640
         # Google Drive storage configuration.
         googleDrive:
           # Enable Google Drive storage.
           enable: false
           # Directory where the password Database file resides.
           directory: "password-manager"
           # Password Database file name.
           passwordDBFile: "password_db"    
        ```
        
    1. Export Environment variable to point the configuration file.
    
        `export PM_CONF_PATH=${PATH_TO_PASSOWRD_MANAGER_YAML}`
1. Follow the instructions to override default configurations using environment variables.
    1. Export required override configuration as environment variable. For example let's assume that you need 
    to override **file permission** of the File storage type. 
        ```$xslt
        export PM_STORAGE_FILE_PERMISSION=0640
        ```   
     Environment variable should have the prefix `PM_` and the hierarchy separation with `_` and keyword **ALL CAPS**.
    
     examples:
    
        `PM_SELECTLISTSIZE=2`
        `PM_ENCRYPTORID=AES`
        `PM_STORAGE_FILE_PASSOWRDDBFILE=password_db`
### Export passwords
If you need to transfer your passwords to a different PC where you have installed password-manager, 
it can be achieved by making a copy of your password database file assuming storage type is **File**. 

Also, you can export passwords to **CVS file** with the following command,
```$xslt
pasword-manager export --csv-file ${PATH_TO_CSV_FILE}
```

Exported CSV file will have the following format([Ex: test/mock-data/data.csv](test/mock-data/data.csv)),
```$xslt
id,username,password,labels
foo@foo.com,foo@foo.com,gijggx3MDxZ,"foo,com"
```
   
### Import Passwords
You can import passwords from a **CSV file** with following command,

```$xslt
pasword-manager import --csv-file ${PATH_TO_CSV_FILE}
```

CSV file should be in the following format([Ex: test/mock-data/data.csv](test/mock-data/data.csv)),
```$xslt
id,username,password,labels
foo@foo.com,foo@foo.com,gijggx3MDxZ,"foo,com"
```
### STORAGE PRECEDENCE
If all the Storage types are enabled, the priority will be given as follow,

* Google drive storage
* Local file storage
 
### SEE ALSO

* [password-manager init](doc/password-manager_init.md)	 - Initialize the Password Manager
* [password-manager add](doc/password-manager_add.md)	 - Add a new password
* [password-manager change](doc/password-manager_change.md)	 - Change a password entry
* [password-manager change-master-password](doc/password-manager_change-master-password.md)	 - Change Master password
* [password-manager generate-password](doc/password-manager_generate-password.md)	 - Generate a secure password
* [password-manager get](doc/password-manager_get.md)	 - Get a password
* [password-manager search](doc/password-manager_search.md)	 - Search Password with ID
* [password-manager import](doc/password-manager_import.md)	 - Import passwords
* [password-manager export](doc/password-manager_export.md)	 - Export password repository to a file
* [password-manager remove](doc/password-manager_remove.md)	 - Remove a password



