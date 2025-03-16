# Simple setup setting

1. Create a `start.sh` script somewhere with the following content:

```bash
if pgrep -x "desktop-spotlight-camera-control" > /dev/null
then
    echo "Killing existing process"
    kill $(pgrep -x "desktop-spotlight-camera-control")
fi

echo "Setting workspace..."
cd /Users/USERNAME/Automations

echo "Loading credentials..."
set -a && source /Users/USERNAME/Automations/CameraControl/.env && set +a

echo "Starting..."
/Users/USERNAME/Automations/CameraControl/desktop-spotlight-camera-control

echo "Ready!"
```

2. Replace `USERNAME` with your username and change the path to the path you used.

3. Save the following `crontab -e`

```bash
* * * * * bash -c 'sh /Users/n.fernandez/Automations/CameraControl/start.sh'
```

You're set!
