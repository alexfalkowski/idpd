# frozen_string_literal: true

module Idpd
  module V1
    class HTTP < Nonnative::HTTPClient
      def create_pipeline(params, opts = {})
        post('/v1/pipeline', params.to_json, opts)
      end
    end
  end
end
