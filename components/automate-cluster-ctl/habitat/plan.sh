#stable channel


pkg_name=automate-ha-cluster-ctl
pkg_description="A package for automate-cluster-ctl for the A2 HA Backend."
pkg_origin=chef
pkg_version="0.1.0"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=("Chef-MLSA")

pkg_deps=(
  core/ruby3_1
  core/aws-cli
  # chef/inspec/4.56.61
  core/bash
  core/coreutils
  core/cacerts
  core/findutils
  core/gawk
  core/gzip
  core/jq-static
  core/openssh
  core/make
  core/curl
  core/terraform
  core/hab/1.6.1243/20241227194506
  core/openssl
  core/rsync
  
#   core/openssh ( core/openssh/7.9p1/20220311132702 core/openssh/7.9p1/20230816095218 )
#   core/openssl ( core/openssl/1.0.2zf/20230816063107 core/openssl/1.0.2zb/20220311111046 )
#   core/ruby30 ( core/ruby30/3.0.3/20220312100602 core/ruby30/3.0.3/20230818185900 )
)

pkg_build_deps=(
  core/gcc
)

pkg_bin_dirs=(bin)

do_before() {
  update_pkg_version
}

do_setup_environment() {
  set_buildtime_env SSL_CERT_FILE "$(pkg_path_for cacerts)/ssl/cert.pem"
}

do_prepare() {
  gem update --system --no-document
  gem install bundler -v "$(grep -A 1 "BUNDLED WITH" $PLAN_CONTEXT/../Gemfile.lock | tail -n 1)"
}

do_build() {
  pushd "$HAB_CACHE_SRC_PATH/$pkg_dirname/" > /dev/null
    if [[ $HAB_DEV != "true" ]]; then
      echo "Cleaning vendor/bundle"
      rm -rf vendor/bundle
    fi
    bundle install --path vendor/bundle
    fix_interpreter "libexec/*" core/coreutils bin/env
  popd
}

do_clean() {
  if [[ $HAB_DEV != "true" ]]; then
    do_default_clean
  else
    # for the dev environment clean up everything but vendor/bundle
    # just clean up the code directories and leave vendor alone
    if [ -d "$HAB_CACHE_SRC_PATH/$pkg_dirname/" ]; then
      pushd "$HAB_CACHE_SRC_PATH/$pkg_dirname/" > /dev/null
        find . -path './vendor' -prune -o -exec rm {} \;
      popd
    fi
  fi
}

do_install() {
  build_line "Copying automate-cluster-ctl files"
  pushd "$HAB_CACHE_SRC_PATH/$pkg_dirname/"
    mkdir -p $pkg_prefix/inspec
    cp -r lib $pkg_prefix/
    cp -r libexec $pkg_prefix/
    cp -r $PLAN_CONTEXT/../../../inspec/automate-backend-opensearch-smoke $pkg_prefix/inspec/
    cp -r $PLAN_CONTEXT/../../../inspec/automate-backend-postgresql-smoke $pkg_prefix/inspec/
    cp -r $PLAN_CONTEXT/../../../inspec/automate-backend-resources $pkg_prefix/inspec/
    cp -r $PLAN_CONTEXT/../../../inspec/automate-frontend-chef-server-smoke $pkg_prefix/inspec/
    cp -r $PLAN_CONTEXT/../../../inspec/automate-frontend-smoke $pkg_prefix/inspec/
    cp -r templates $pkg_prefix/
    cp -r vendor $pkg_prefix/
  popd
  mkdir -p "$pkg_prefix/bin"
  wrap_bin "automate-cluster-ctl"
}

do_unpack() {
  mkdir -pv "$HAB_CACHE_SRC_PATH/$pkg_dirname"
  cp -RT "$PLAN_CONTEXT"/.. "$HAB_CACHE_SRC_PATH/$pkg_dirname/"
}

do_setup_environment() {
  export GEM_HOME="$pkg_prefix/vendor/bundle/ruby/3.1.0"
  export GEM_PATH="$GEM_HOME"

  set_runtime_env GEM_HOME "$GEM_HOME"
  set_buildtime_env GEM_HOME "$GEM_HOME"
  push_runtime_env GEM_PATH "$GEM_PATH"
  push_buildtime_env GEM_PATH "$GEM_PATH"
}

# Need to wrap the binary to ensure GEM_HOME/GEM_PATH is correct
wrap_bin() {
  local bin="$pkg_prefix/bin/$1"
  local real_bin="$pkg_prefix/libexec/$1"

  build_line "Creating wrapper $bin for $real_bin"
  cat <<EOF > "$bin"
#!$(pkg_path_for core/bash)/bin/bash
set -e

source $pkg_prefix/RUNTIME_ENVIRONMENT
export GEM_PATH GEM_HOME PATH
export AUTOMATE_CLUSTER_VERSION="${pkg_version}"

exec $real_bin \$@
EOF
  chmod -v 755 "$bin"
}

do_strip() {
  return 0
}


