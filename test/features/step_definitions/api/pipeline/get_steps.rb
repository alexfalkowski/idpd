# frozen_string_literal: true

When('we try to get an nonexistent pipeline') do
  @response = Idpd.http.get_pipeline('123456', Idpd.options)
end

When('we get an invalid pipeline') do
  @response = Idpd.http.get_pipeline(0, Idpd.options)
end

When('we get the created simple pipeline') do
  @id = JSON.parse(@response.body)['pipeline']['id']
  @response = Idpd.http.get_pipeline(@id, Idpd.options)
end

Then('we should have a simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(@id)
end

Then('we should have a not found request') do
  expect(@response.code).to eq(404)
end
