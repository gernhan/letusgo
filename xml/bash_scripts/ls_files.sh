#!/usr/bin/expect -f

# Replace the following variables with your actual SFTP server details
set SFTP_HOST "34.89.153.28"
set SFTP_PORT "22"
set SFTP_USER "compax"
set SFTP_PASSWORD "compax"

# Connect to the SFTP server
spawn sftp -oPort=$SFTP_PORT $SFTP_USER@$SFTP_HOST

# Expect and respond to password prompt
expect "password:"
send "$SFTP_PASSWORD\r"

# Execute the 'ls' command on the SFTP server and store the result in a variable
expect "sftp>"
send "ls -1 upload/test6\r"
expect "sftp>"
set sftp_result $expect_out(buffer)

puts "$sftp_result"

# Loop through each line in the result and execute 'ls' for each item
set lines [split $sftp_result "\n"]
foreach line $lines {
    if {[string length $line] > 0} {
        # Check if the line is a folder (no spaces in the name)
        if {! [string match "* *" $line]} {
            # Change into the folder
            send "cd $line\r"
            expect "sftp>"
            send "ls\r"
            expect "sftp>"
            set files $expect_out(buffer)

            set file_lines [split $files "\n"]
            # Loop through each line (file) and remove it
            foreach file_line $file_lines {
                if {[string length $file_line] > 0} {
                    # Remove the file
                    send "cd $file_line\r"
                    send "ls -1 $file_line"
                    send "cd\r"
                }
                send "cd\r"
            }
        }
    }
}

# Quit SFTP session
send "quit\r"

# End the script
expect eof

