chown -R dockeruser:dockerusergroup /app
chmod -R g+rwx /app
chown -R dockeruser:dockerusergroup /usr/bin/node && chown -R dockeruser /usr/bin/npm
