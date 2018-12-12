FROM ruby:2.5.3
RUN useradd -ms /bin/bash app
ENV LANG C.UTF-8
ENV RAILS_ENV production
ENV SECRET_KEY_BASE test
COPY . /home/app
RUN chown -R app:app /home/app
USER app
WORKDIR /home/app
RUN bundle install --deployment --without development test --jobs=3
EXPOSE 3000
ENTRYPOINT ["bundle", "exec", "puma"]
