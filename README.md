# Garmin

Personal set up to import Garmin data from my watch and visualize it in Grafana.

This project is not intended to work more generally.

Related art:

- https://github.com/tcgoetz/GarminDB/
- https://github.com/garmin/fit-python-sdk/

## Sleep data

Definitions:

> **Note:** fields are 1-indexed.
> **Note:** cells with a `(?)` are best guesses.

| Local number | Message         | Field index | Field value               | Units |
|--------------|-----------------|-------------|---------------------------|-------|
| 0            | File ID         | 1           | Serial number             |       |
| 0            | File ID         | 2           | Time created              |       |
| 0            | File ID         | 3           | Unknown                   |       |
| 0            | File ID         | 4           | Manufacturer              |       |
| 0            | File ID         | 5           | Product                   |       |
| 0            | File ID         | 6           | Number                    |       |
| 0            | File ID         | 7           | Type                      |       |
| 1            | File creator    | 1           | Unknown                   |       |
| 1            | File creator    | 2           | Software version          |       |
| 1            | File creator    | 3           | Hardware version          |       |
| 2            | Device info     | 1           | Timestamp                 |       |
| 2            | Device info     | 2           | Serial number             |       |
| 2            | Device info     | 3           | Cumulative operating time |       |
| 2            | Device info     | 4           | Unknown                   |       |
| 2            | Device info     | 5           | Unknown                   |       |
| 2            | Device info     | 6           | Unknown                   |       |
| 2            | Device info     | 7           | Unknown                   |       |
| 2            | Device info     | 8           | Unknown                   |       |
| 2            | Device info     | 9           | Unknown                   |       |
| 2            | Device info     | 10          | Manufacturer              |       |
| 2            | Device info     | 11          | Product                   |       |
| 2            | Device info     | 12          | Software version          |       |
| 2            | Device info     | 13          | Battery voltage           |       |
| 2            | Device info     | 14          | Unknown                   |       |
| 2            | Device info     | 15          | ANT device number         |       |
| 2            | Device info     | 16          | Device index              |       |
| 2            | Device info     | 17          | Device type               |       |
| 2            | Device info     | 18          | Hardware version          |       |
| 2            | Device info     | 19          | Unknown                   |       |
| 2            | Device info     | 20          | Battery status            |       |
| 2            | Device info     | 21          | Sensor position           |       |
| 2            | Device info     | 22          | ANT transmission type     |       |
| 2            | Device info     | 23          | ANT network               |       |
| 2            | Device info     | 24          | Unknown                   |       |
| 2            | Device info     | 25          | Source type               |       |
| 2            | Device info     | 26          | Unknown                   |       |
| 2            | Device info     | 27          | Unknown                   |       |
| 2            | Device info     | 28          | Battery level             |       |
| 3            | Unknown         | 1           | Unknown                   |       |
| 3            | Unknown         | 2           | Unknown                   |       |
| 3            | Unknown         | 3           | Unknown                   |       |
| 3            | Unknown         | 4           | Unknown                   |       |
| 3            | Unknown         | 5           | Unknown                   |       |
| 3            | Unknown         | 6           | Unknown                   |       |
| 4            | Event           | 1           | Timestamp                 |       |
| 4            | Event           | 2           | Data                      |       |
| 4            | Event           | 3           | Event                     |       |
| 4            | Event           | 4           | Event type                |       |
| 4            | Event           | 5           | Event group               |       |
| 4            | Event           | 6           | Unknown                   |       |
| 4            | Event           | 7           | Unknown                   |       |
| 5            | Unknown         | 1           | Unknown                   |       |
| 5            | Unknown         | 2           | Unknown                   |       |
| 5            | Unknown         | 3           | Unknown                   |       |
| 5            | Unknown         | 4           | Unknown                   |       |
| 5            | Unknown         | 5           | Unknown                   |       |
| 5            | Unknown         | 6           | Unknown                   |       |
| 5            | Unknown         | 7           | Unknown                   |       |
| 5            | Unknown         | 8           | Unknown                   |       |
| 5            | Unknown         | 9           | Unknown                   |       |
| 5            | Unknown         | 10          | Unknown                   |       |
| 5            | Unknown         | 11          | Unknown                   |       |
| 5            | Unknown         | 12          | Unknown                   |       |
| 5            | Unknown         | 13          | Unknown                   |       |
| 5            | Unknown         | 14          | Unknown                   |       |
| 5            | Unknown         | 15          | Unknown                   |       |
| 5            | Unknown         | 16          | Unknown                   |       |
| 6            | Sleep event (?) | 1           | End timestamp (?)         |       |
| 6            | Sleep event (?) | 2           | Sleep level (?)           |       |
| 7            | Unknown         | 1           | Unknown                   |       |
| 7            | Unknown         | 2           | Unknown                   |       |
| 7            | Unknown         | 3           | Unknown                   |       |

## Running data

Definitions:

> **Note:** fields are 1-indexed.
> **Note:** cells with a `(?)` are best guesses.

| Local number | Message | Field index | Field value                 | Units       |
|--------------|---------|-------------|-----------------------------|-------------|
| 7            | Session | 1           | Timestamp                   | s           |
| 7            | Session | 2           | Start time                  |             |
| 7            | Session | 3           | Start position latitude     | semicircles |
| 7            | Session | 4           | Start position longitude    | semicircles |
| 7            | Session | 5           | Total elapsed time          | s           |
| 7            | Session | 6           | Total timer time            | s           |
| 7            | Session | 7           | Total distance              | m           |
| 7            | Session | 8           | Total strides               | cycles      |
| 7            | Session | 9           | North-east corner latitude  | semicircles |
| 7            | Session | 10          | North-east corner longitude | semicircles |
| 7            | Session | 11          | South-west corner latitude  | semicircles |
| 7            | Session | 12          | South-west corner longitude | semicircles |
| 7            | Session | 13          | Unknown                     |             |
| 7            | Session | 14          | Unknown                     |             |

## Bouldering data
