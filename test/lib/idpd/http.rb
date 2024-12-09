# frozen_string_literal: true

module Idpd
  class HTTP < Nonnative::HTTPClient
    def create_pipeline(pipeline, opts = {})
      post('/pipelines', pipeline, opts)
    end

    def get_pipeline(id, opts = {})
      get("/pipelines/#{id}", opts)
    end

    def update_pipeline(id, pipeline, opts = {})
      put("/pipelines/#{id}", pipeline, opts)
    end

    def delete_pipeline(id, opts = {})
      delete("/pipelines/#{id}", opts)
    end

    def trigger_pipeline(id, opts = {})
      post("/pipelines/#{id}/trigger", '', opts)
    end
  end
end
