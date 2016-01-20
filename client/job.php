<?php

class Job {
    public $name;
    public $url;
    function __construct($name, $url)
    {
        $this->name = $name;
        $this->url = $url;
    }

    public function getDescription() {
        return 'jkljlkjlk ' . $this->name . 'jkjkljkljkljlk' . $this->url;
    }
}


$job = new Job('jan', 'jkljkljkl');

echo $job->getDescription();
