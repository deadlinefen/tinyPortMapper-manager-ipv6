# config.toml

`config_description_en.md` has been translated using Baidu's Wenxin Yiyan (文心一言) translation tool.

The configuration file consists of four sections: mapper, resolution, jobs, and log. Among them, mapper is related to tinyPortMapper configuration, resolution is related to domain name resolution configuration, jobs is for port mapping task configuration, and log is for logging configuration.

## mapper

1. **bin**
   The path where the tinyPortMapper executable file is located.
2. **file-directory**
   The directory used to save the log files generated by tinyPortMapper. Since there may be multiple mapping relationships, this field specifies the directory for logs, and the file naming format is `fromPort-toIp:toPort.log`.

## resolution

1. domain
   The domain name that needs to be resolved. The current version only supports one domain name.
2. dns
   A list of DNS server addresses used for resolution. To prevent frequent access and potential blacklisting of nodes, it is recommended to provide multiple entries. The manager will iterate through the list of DNS servers to perform domain name resolution.
3. ttl
   Every ttl seconds, the manager selects a DNS server from the list to perform domain name resolution.

## jobs

Multiple jobs can be specified, with each job corresponding to a port mapping task. Each task is identified by `[jobs.job_name]`.

1. from-port
   Corresponds to the local-port of tinyPortMapper.
2. to-ip
   Corresponds to the remote-ip of tinyPortMapper.
3. to-port
   Corresponds to the remote-port of tinyPortMapper.
4. type
   Corresponds to the mapping type of tinyPortMapper. `"t"` represents TCP, `"u"` represents UDP, and `tu` represents both TCP and UDP.

## log

1. level
   The logging level, with debug > info > warn > error > fatal > panic. For daily use, it is recommended to use info.
2. path
   The path where logs are stored.
3. to-stdout-only
   Whether to only output logs to stdout. For daily use, it is recommended to set this to false.
4. also-to-stderr
   Whether to simultaneously output logs to stderr.
