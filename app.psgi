use Kelp::Less;
use Try::Tiny;

get '/put/:x/:y' => sub {
    my ( $self, $x, $y ) = @_;
    { x => $x, y => $y };
};

get '/get' => sub {
    my $self   = shift;
    my $result = "OK";
    try {
        $self->json->decode( param('p') );
    }
    catch {
        $result = "ERROR: $_";
    };
    $result;
};

run;
