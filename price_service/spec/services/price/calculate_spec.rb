require 'rails_helper'

describe Price::Calculate do
  subject { described_class.run!(params) }

  let(:params) do
    {
      start_point: {
        lat: 55.729060,
        lng: 37.622691
      },
      end_point: {
        lat: 55.808116,
        lng: 37.581609
      }
    }
  end

  before do
    result = {
      distance: 12135.455,
      time: 1284809
    }

    allow(Direction::Api).to receive(:calculate).and_return(result)
  end

  it 'returns price' do
    expect(subject).to eq(4775.to_d)
  end
end
