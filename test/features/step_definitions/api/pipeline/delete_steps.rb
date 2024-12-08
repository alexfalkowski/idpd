# frozen_string_literal: true

When('we delete the simple pipeline') do
  @id = JSON.parse(@response.body)['pipeline']['id'].to_i
  @response = Idpd::V1.server.delete_pipeline(@id, opts)
end

When('we try to delete an nonexistent pipeline') do
  @response = Idpd::V1.server.delete_pipeline(10, opts)
end

When('we try to delete with an invalid id') do
  @response = Idpd::V1.server.delete_pipeline(0, opts)
end

Then('we should have deleted simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(@id)
end
