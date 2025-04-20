import json
import random
import os
from datetime import datetime, timedelta

# Sample data for CloudWatch logs
log_levels = ["INFO", "ERROR", "WARNING", "DEBUG"]
regions = ["us-east-1", "us-west-2", "eu-central-1", "ap-southeast-1"]
log_groups = ["/aws/lambda/auth-service", "/aws/ec2/system-metrics", "/aws/rds/db-logs", "/aws/ecs/service-logs"]
log_streams = {
    "/aws/lambda/auth-service": ["2025/04/20/[$LATEST]abcd1234", "2025/04/20/[$LATEST]efgh5678"],
    "/aws/ec2/system-metrics": ["i-0123456789abcdef0", "i-0fedcba9876543210"],
    "/aws/rds/db-logs": ["rds-instance-01-log", "rds-instance-02-log"],
    "/aws/ecs/service-logs": ["ecs-service-A-task1234", "ecs-service-B-task5678"]
}

def generate_cloudwatch_log_entry():
    log_group = random.choice(log_groups)
    log_stream = random.choice(log_streams[log_group])
    log_level = random.choice(log_levels)
    region = random.choice(regions)
    timestamp = (datetime.utcnow() - timedelta(seconds=random.randint(0, 1000000))).isoformat() + "Z"
    message = f"{log_level}: Simulated CloudWatch log for {log_group} in {region}"
    ingestion_time = int((datetime.utcnow() - timedelta(seconds=random.randint(0, 1000000))).timestamp() * 1000)
    
    return {
        "timestamp": timestamp,
        "message": message,
        "logGroup": log_group,
        "logStream": log_stream,
        "region": region,
        "ingestionTime": ingestion_time,
        "eventId": ''.join(random.choices("abcdef0123456789", k=32))
    }

folder_path = "../Data"
file_path = os.path.join(folder_path, "cloudwatch_logs.jsonl")

# Create folder if it doesn't exist
os.makedirs(folder_path, exist_ok=True)

# Only generate and write logs if the file does not exist
if not os.path.exists(file_path):
    aws_logs = [json.dumps(generate_cloudwatch_log_entry()) for _ in range(1_000_000)]
    with open(file_path, "w") as f:
        f.write("\n".join(aws_logs))
    print("Plik został utworzony i zapisany.")
else:
    print("Plik już istnieje — nic nie zostało nadpisane.")