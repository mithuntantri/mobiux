# Sales Report

Sales Report generator for the data from an ice cream parlour. This program is written in Golang

## Prerequisites

```
Golang : v1.13 and above
export GOPATH = $HOME/mobuiux
```
## Usage

### Golang Applications

Suitable for applications that have their concurrent processing and serverless 
lambda architecture or docker applications.

<details><summary><b>Show instructions</b></summary>

1. Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go and Add /usr/local/go/bin to the PATH environment variable.:

    ```sh
    $ tar -C /usr/local -xzf go1.15.4.linux-amd64.tar.gz
    $ export PATH=$PATH:/usr/local/go/bin
    $ export GOPATH = $HOME/mobuiux
    ```

2. Build the binary `salestrack` run it with `salestrack` command:

    ```sh
    cd $HOME/mobiux/src/salestrack
    go install
    salestrack
    ```
3. Here's how it generated the result

	```
	---------------------------------------------------------------------------------------------------------------------------------
	| YEAR	| MONTH		| SALES		| Most Popular Item				| Most Revenue Generating Item		|
	---------------------------------------------------------------------------------------------------------------------------------
	| 2019	|    January	| 1421330	| Death by Chocolate (Min:1, Max:5, Avg:3)	| Hot Chocolate Fudge (Revenue:320760)	|
	| 2019	|   February	| 1422350	| Vanilla Single Scoop (Min:1, Max:5, Avg:3)	| Hot Chocolate Fudge (Revenue:316320)	|
	| 2019	|      March	| 1739590	| Pista Double Scoop (Min:1, Max:5, Avg:3)	| Hot Chocolate Fudge (Revenue:372000)	|
	---------------------------------------------------------------------------------------------------------------------------------
	TOTAL SALES :::::::::::::::  4583270
	```