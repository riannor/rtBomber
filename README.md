# rtBomber
Simple and flexible UDP flood DoS test tool. 

## Usage

Tool can be used in interactive or package mode

### Interactive mode:

Used for one site. Tool ask parameters after start:
> Domain:

Site domain should be set (without "http:" or "www" prefix) 

> iterations (*100000): 

Count of iteration: if you set value 100, tool make flood with 1 million packets 
to given site

> threads:

All packets count will be splitted to given parallel threads 

After setup all settings test will start. Progress (in percent of all iterations)
will be displayed. After all iteration finish, tool will return to interactive mode
and will ask new domain and all other settings.

### Package mode:

Can be used for test several sites one-by-one. Make text file "targets.txt" and 
put it near to tool binary. All domains should be put to file in separate lines.
If file will be found after start, tool will ask other settings in interactive 
mode and start test all sites with same settings. Progress for each site will be
displayed. After last site test will finish tool will be stopped.

### Default settings:

port: 443

UDP packet size:  65507