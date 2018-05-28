# -*- mode: ruby -*-
# vi: set ft=ruby :

$num_instances = 2


Vagrant.configure("2") do |config|



  # config.vm.network "forwarded_port", guest: 80, host: 8080
  
    (1..$num_instances).each do |i|
      config.vm.define vm_name = "%s-%02d" % ["eyeswarm", i] do |config|
      config.vm.hostname = vm_name
      config.vm.box = "ubuntu/xenial64"
      config.vm.provider :virtualbox do |vb|
        vb.memory = "1024"
        vb.cpus = 2
        vb.customize ["modifyvm", :id, "--cpuexecutioncap", "100"]
      end
      ip = "172.17.8.#{i+10}"
      config.vm.network :private_network, ip: ip
      config.vm.provision :file, :source => "deploy.sh", :destination => "/tmp/vagrantfile-user-data"
    end
  
  # config.vm.provision "shell", inline: <<-SHELL
  #   apt-get update
  #   apt-get install -y apache2
  # SHELL
end
end
