<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Questionpools_testing extends Model
{
    protected $fillable = [
        'qpid', 'input', 'output',
    ];
}
