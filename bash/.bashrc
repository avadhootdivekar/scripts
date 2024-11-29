
# AD : Infinite history 
# Eternal bash history.
# ---------------------
# Undocumented feature which sets the size to "unlimited".
# http://stackoverflow.com/questions/9457233/unlimited-bash-history
export HISTFILESIZE=
export HISTSIZE=
export HISTTIMEFORMAT="[%F %T] "
# Change the file location because certain bash sessions truncate .bash_history file upon close.
# http://superuser.com/questions/575479/bash-history-truncated-to-500-lines-on-each-login
export HISTFILE=~/.bash_eternal_history
# Force prompt to write history after every command.
# http://superuser.com/questions/20900/bash-history-loss
PROMPT_COMMAND="history -a; $PROMPT_COMMAND"
# AD Infinite history end. 
# AD 
export WHOME=/mnt/c/Users/avadh/
export WCODE=/mnt/c/Users/avadh/code/
export CODE=/home/avadhoot/mounted/code
export PATH=$PATH:/usr/local/go/bin
export PATH_ATMOS_KUBERNETES=$CODE
alias k=kubectl
source <(kubectl completion bash)
source ~/.helper.sh
alias mycmd='source ~/.helper.sh'