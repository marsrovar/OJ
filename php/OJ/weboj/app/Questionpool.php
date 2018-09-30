<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Questionpool extends Model
{
    protected $fillable = [
        'num', 'title', 'content', 'code_function','inputtype',
    ];
}
