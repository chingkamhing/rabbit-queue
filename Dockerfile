FROM rabbitmq:3.7-management-alpine

RUN rabbitmq-plugins enable --offline rabbitmq_mqtt
RUN rabbitmq-plugins enable --offline rabbitmq_auth_backend_http
RUN rabbitmq-plugins enable --offline rabbitmq_auth_backend_cache

# COPY ./rabbitmq.conf /etc/rabbitmq/

#rabbitmq-amqp
EXPOSE 5672

RUN rabbitmq-plugins enable --offline rabbitmq_web_mqtt

#rabbitmq-mqtt-over-ws
#expose to public as 443 https
EXPOSE 15675
