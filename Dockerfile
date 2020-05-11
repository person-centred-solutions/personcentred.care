FROM ruby:2-slim

WORKDIR /app

RUN apt-get update && \
    apt-get install -y git build-essential curl gnupg
# Install Node JS
RUN curl -sL https://deb.nodesource.com/setup_12.x  | bash -
RUN apt-get -y install nodejs

# Setup ruby and local plugins
ADD Gemfile Gemfile.lock /app/
ADD .git /app/.git
ADD jekyll-seo-tag /app/jekyll-seo-tag

RUN ls
RUN bundle install

CMD bundle exec jekyll serve
