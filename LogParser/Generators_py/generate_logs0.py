# Re-import necessary libraries after kernel reset
import json
import os
import random
from datetime import datetime, timedelta

# Sample AWS log types
aws_services = ["EC2", "S3", "Lambda", "CloudWatch", "IAM", "RDS", "ECS", "DynamoDB"]
log_levels = ["INFO", "ERROR", "WARNING", "DEBUG"]
actions = {
    "EC2": ["StartInstances", "StopInstances", "RebootInstances"],
    "S3": ["PutObject", "GetObject", "DeleteObject"],
    "Lambda": ["Invoke", "CreateFunction", "DeleteFunction"],
    "CloudWatch": ["PutMetricData", "GetMetricData", "DescribeAlarms"],
    "IAM": ["CreateUser", "DeleteUser", "AttachRolePolicy"],
    "RDS": ["CreateDBInstance", "DeleteDBInstance", "ModifyDBInstance"],
    "ECS": ["RunTask", "StopTask", "DescribeTasks"],
    "DynamoDB": ["PutItem", "GetItem", "DeleteItem"]
}

# Generate random AWS-style log entry
def generate_aws_log_entry():
    service = random.choice(aws_services)
    action = random.choice(actions[service])
    log_level = random.choice(log_levels)
    timestamp = (datetime.utcnow() - timedelta(seconds=random.randint(0, 1000000))).isoformat() + "Z"
    return {
        "timestamp": timestamp,
        "logLevel": log_level,
        "service": service,
        "action": action,
        "message": f"{service} {action} executed with log level {log_level}",
        "requestId": ''.join(random.choices("abcdef0123456789", k=16))
    }

folder_path = "../Data"
file_path = os.path.join(folder_path, "aws_logs.jsonl")

# Create folder if it doesn't exist
os.makedirs(folder_path, exist_ok=True)

# Only generate and write logs if the file does not exist
if not os.path.exists(file_path):
    aws_logs = [json.dumps(generate_aws_log_entry()) for _ in range(1_000_000)]
    with open(file_path, "w") as f:
        f.write("\n".join(aws_logs))
    print("Plik został utworzony i zapisany.")
else:
    print("Plik już istnieje — nic nie zostało nadpisane.")