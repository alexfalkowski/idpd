# frozen_string_literal: true

lib = File.expand_path('../../lib', __dir__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require 'nonnative'
require 'idpd'

def opts
  {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }
end
