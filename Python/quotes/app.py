import os
import boto3
import grpc
import redis

from flask import Flask, jsonify, make_response, request

BUCKET_NAME = 'serverless-machine-learning'

dynamodb_client = boto3.client('dynamodb')

grpc_host = "localhost:50051"

channel = grpc.insecure_channel(grpc_host)

S3 = boto3.client('s3', region_name='eu-central-1')

r = redis.Redis(host='localhost', port=6379, db=0)

# response = S3.get_object(Bucket=BUCKET_NAME, Key=key)

app = Flask(__name__)

if os.environ.get('IS_OFFLINE'):
    dynamodb_client = boto3.client(
        'dynamodb', 
        region_name='localhost', 
        endpoint_url='http://localhost:8000'
    )

USERS_TABLE = os.environ['USERS_TABLE']

@app.route("/")
def hello_from_root():
    return jsonify(message='Do not wanna attack my server !')

@app.route('/users/<string:user_id>')
def get_user(user_id):
    result = dynamodb_client.get_item(
        TableName=USERS_TABLE, Key={'userId': {'S': user_id}}
    )
    item = result.get('Item')
    if not item:
        return jsonify({'error': 'Could not find user with provided "userId"'}), 404

    return jsonify(
        {'userId': item.get('userId').get('S'), 'name': item.get('name').get('S')}
    )

@app.route('/users', methods=['POST'])
def create_user():
    user_id = request.json.get('userId')
    name = request.json.get('name')
    if not user_id or not name:
        return jsonify({'error': 'Please provide both "userId" and "name"'}), 400

    dynamodb_client.put_item(
        TableName=USERS_TABLE, Item={'userId': {'S': user_id}, 'name': {'S': name}}
    )

    return jsonify({'userId': user_id, 'name': name})

@app.errorhandler(404)
def resource_not_found(e):
    return make_response(jsonify(error='Not found!'), 404)