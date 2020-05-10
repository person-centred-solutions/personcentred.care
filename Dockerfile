FROM ruby:2-slim

WORKDIR /app

RUN mkdir /prepenv
ADD Gemfile Gemfile.lock /prepenv

RUN apt-get update && \
    apt-get install -y git build-essential
RUN cd /prepenv && bundle install

CMD bundle exec jekyll serve
