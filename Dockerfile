FROM ruby:2-slim

WORKDIR /app

ADD Gemfile Gemfile.lock /app/

RUN apt-get update && \
    apt-get install -y git build-essential
RUN bundle install

CMD bundle exec jekyll serve
