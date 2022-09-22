#!/bin/bash

enc=`urlencode $3`
encuser=`urlencode $2`

cp /account.config ~
sed -i s/DIREKTIVSERVER/$1/ ~/account.config 
sed -i s/DIREKTIVUSER/${encuser}/ ~/account.config 
sed -i s/DIREKTIVPWD/${enc}/ ~/account.config 
