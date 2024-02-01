import logging
import platform
import time

import influxdb_client
import psutil
from influxdb_client import Point
from influxdb_client.client.write_api import SYNCHRONOUS

# This is the token at zhefuz.link:8086 influxdb

org = "zhefuz.link"
url = "http://zhefuz.link:8086"
bucket = "default_bucket"
write_client = influxdb_client.InfluxDBClient(url=url, username="root", password="pA1zsWO10FDsBd", org=org)
write_api = write_client.write_api(write_options=SYNCHRONOUS)


class Record:
	def __init__(self, hostname, cpu_usage, memory_usage, disk_usage):
		self.hostname = hostname
		self.cpu_usage = cpu_usage
		self.memory_usage = memory_usage
		self.disk_usage = disk_usage


def collect_system_stats():
	hostname = platform.uname()[1]
	cpu_usage = psutil.cpu_percent(interval=None, percpu=False)

	memory_usage = psutil.virtual_memory().percent
	disk_usage = psutil.disk_usage("/").percent
	return Record(hostname, cpu_usage, memory_usage, disk_usage)


def write():
	current = collect_system_stats()
	t = time.time_ns()
	point = (
		Point("cpu_usage")
		.time(t)
		.tag("hostname", current.hostname)
		.field("percent", current.cpu_usage)
	)
	try:
		write_api.write(bucket=bucket, org=org, record=point)
		logging.debug(
			f"Collecting...\n"
			f"Hostname: {current.hostname}\n"
			f"CPU Usage: {current.cpu_usage}\n")
	except:
		logging.critical("Connection failed")
	point = (
		Point("memory_usage")
		.time(t)
		.tag("hostname", current.hostname)
		.field("percent", current.memory_usage)
	)
	try:
		write_api.write(bucket=bucket, org=org, record=point)
		logging.debug(
			f"Memory Usage: {current.memory_usage}\n"
		)
	except:
		logging.critical("Connection failed")
	point = (
		Point("disk_usage")
		.time(t)
		.tag("hostname", current.hostname)
		.field("percent", current.disk_usage)
	)
	try:
		write_api.write(bucket=bucket, org=org, record=point)
		logging.debug(
			f"Disk Usage: {current.disk_usage}\n"
		)
	except:
		logging.critical("Connection failed")


if __name__ == '__main__':
	logging.basicConfig(format='%(asctime)s - %(pathname)s[line:%(lineno)d]\n%(levelname)s: %(message)s',
						level=logging.DEBUG)
	# connect()
	while 1:
		write()
		time.sleep(2)
