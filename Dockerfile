FROM ruby:2-slim

WORKDIR /app

ADD Gemfile Gemfile.lock /app/

RUN apt-get update && \
    apt-get install -y git build-essential curl gnupg
RUN curl -sL https://deb.nodesource.com/setup_12.x  | bash -
RUN apt-get -y install nodejs
RUN bundle install

CMD bundle exec jekyll serve
