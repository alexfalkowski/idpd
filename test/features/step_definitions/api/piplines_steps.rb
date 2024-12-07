# frozen_string_literal: true

def opts
  {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Idpd.token),
    read_timeout: 10, open_timeout: 10
  }
end

When('we create a simple pipeline') do
  pipeline = {
    pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.create_pipeline(pipeline, opts)
end

When('we try to create a invalid pipeline of kind {string}') do |kind|
  cases = {
    'invalid pipeline name' => { pipeline: {} },
    'missing jobs' => { pipeline: { 'name' => 'test' } },
    'invalid job name' => { pipeline: { 'name' => 'test', 'jobs' => [{ 'steps' => %w[test test2] }] } },
    'missing steps' => { pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test' }] } }
  }

  @response = Idpd::V1.server.create_pipeline(cases[kind], opts)
end

When('we try to create a pipeline with being authorized') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  pipeline = {
    pipeline: { 'name' => 'test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.create_pipeline(pipeline, opts)
end

When('we get the created simple pipeline') do
  @response = Idpd::V1.server.get_pipeline(2, opts)
end

When('we try to get an nonexistent pipeline') do
  @response = Idpd::V1.server.get_pipeline(10, opts)
end

When('we get an invalid pipeline') do
  @response = Idpd::V1.server.get_pipeline(0, opts)
end

When('we update the simple pipeline') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(3, pipeline, opts)
end

When('we try to update an nonexistent pipeline') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(10, pipeline, opts)
end

When('we try to update an invalid pipeline') do
  pipeline = {
    pipeline: { 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(3, pipeline, opts)
end

When('we try to update with an invalid id') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(0, pipeline, opts)
end

When('we delete the simple pipeline') do
  @response = Idpd::V1.server.delete_pipeline(4, opts)
end

When('we try to delete an nonexistent pipeline') do
  @response = Idpd::V1.server.delete_pipeline(10, opts)
end

When('we try to delete with an invalid id') do
  @response = Idpd::V1.server.delete_pipeline(0, opts)
end

Then('we should have a created pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(1)
end

Then('we should have bad request') do
  expect(@response.code).to eq(400)
end

Then('we should have a unauthorized request') do
  expect(@response.code).to eq(401)
end

Then('we should have a simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(2)
end

Then('we should have a not found request') do
  expect(@response.code).to eq(404)
end

Then('we should have an updated simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['name']).to eq('update-test')
end

Then('we should have deleted simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(4)
end
