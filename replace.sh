#!/bin/bash

enc=`urlencode $3`

cp /account.config ~
sed -i s/DIREKTIVSERVER/$1/ ~/account.config 
sed -i s/DIREKTIVUSER/$2/ ~/account.config 
sed -i s/DIREKTIVPWD/${enc}/ ~/account.config 
